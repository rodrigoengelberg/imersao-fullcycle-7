import { PixKeyKind } from 'src/models/pix-key.model';

export class PixKeyDto {
  readonly key: string;
  readonly kind: PixKeyKind;
}
