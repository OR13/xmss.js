import { sum } from '../src/index';

describe('sum', () => {
  it('adds two numbers together', async () => {
    expect(await sum(1, 1)).toEqual(2);
  });
});
