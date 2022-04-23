import { Controller, Get, Param } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { BankAccount } from 'src/models/bank-account.model';
import { Repository } from 'typeorm';

@Controller('bank-accounts')
export class BankAccountController {
  constructor(
    @InjectRepository(BankAccount)
    private bankAccountRepo: Repository<BankAccount>,
  ) {}

  @Get()
  index() {
    return this.bankAccountRepo.find();
  }

  @Get(':bankAccountId')
  show(@Param('bankAccountId') bankAccount: string) {
    return this.bankAccountRepo.findOneOrFail(bankAccount);
  }
}
