import { init } from './main';

export const sum = async (a: number, b: number) => {
  const wasm = await init();
  return wasm.add(a, b);
};
