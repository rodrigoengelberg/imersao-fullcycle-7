import {
  Body,
  Controller,
  Get,
  Inject,
  InternalServerErrorException,
  Param,
  ParseUUIDPipe,
  Post,
  UnprocessableEntityException,
  ValidationPipe,
} from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { InjectRepository } from '@nestjs/typeorm';
import { PixKeyDto } from 'src/dto/pix-key.dto';
import { PixService } from 'src/grpc-types/pix-services.grpc';
import { BankAccount } from 'src/models/bank-account.model';
import { PixKey } from 'src/models/pix-key.model';
import { Repository } from 'typeorm';

@Controller('bank-accounts/:bankAccountId/pix-keys')
export class PixKeyController {
  constructor(
    @InjectRepository(PixKey)
    private pixKeyRepo: Repository<PixKey>,
    @InjectRepository(BankAccount)
    private bankAccountRepo: Repository<BankAccount>,
    @Inject('CODEPIX_PACKAGE')
    private client: ClientGrpc,
  ) {}
  @Get()
  index(
    @Param('bankAccountId', new ParseUUIDPipe({ version: '4' }))
    bankAccountId: string,
  ) {
    return this.pixKeyRepo.find({
      where: {
        bank_account_id: bankAccountId,
      },
      order: {
        created_at: 'DESC',
      },
    });
  }

  @Post()
  async store(
    @Param('bankAccountId', new ParseUUIDPipe({ version: '4' }))
    bankAccoutId: string,
    @Body(new ValidationPipe({ errorHttpStatusCode: 422 }))
    body: PixKeyDto,
  ) {
    await this.bankAccountRepo.findOneOrFail(bankAccoutId);

    const pixService: PixService = this.client.getService('PixService');
    const notFound = await this.checkPixKeyNotFound(body);

    if (!notFound) {
      throw new UnprocessableEntityException('PixKey already exists');
    }

    const createdPixKey = await pixService
      .registerPixKey({
        ...body,
        accountId: bankAccoutId,
      })
      .toPromise();

    if (createdPixKey.error) {
      throw new InternalServerErrorException(createdPixKey.error);
    }

    const pixKey = this.pixKeyRepo.create({
      id: createdPixKey.id,
      bank_account_id: bankAccoutId,
      ...body,
    });

    return await this.pixKeyRepo.save(pixKey);
  }

  async checkPixKeyNotFound(params: { key: string; kind: string }) {
    const pixService: PixService = this.client.getService('PixService');
    try {
      await pixService.find(params).toPromise();
      return false;
    } catch (e) {
      if (e.details === 'no key was found') {
        return true;
      }

      throw new InternalServerErrorException('Server not available');
    }
  }

  @Get('exists')
  exists() {}
}
