import { l } from '@/languages/language';
import { isNotEmpty } from './form-validation';

describe('isNotEmpty', () => {
  it('returns error message when the value is empty', () => {
    expect(isNotEmpty('')).toBe(l.comn.form.thisFieldIsRequired);
  });

  it('returns null when the value is not empty', () => {
    expect(isNotEmpty('hello')).toBeNull();
    expect(isNotEmpty(' ')).toBeNull();
  });
});
