import { init } from './main';

export const generate = async () => {
  const wasm = await init();
  const request = JSON.stringify({ command: 'generate' });
  const response = wasm.handleRequest(request);
  return JSON.parse(response);
};

export const sign = async (message: string, jwk: any) => {
  const wasm = await init();
  const request = JSON.stringify({
    command: 'sign',
    message,
    jwk: JSON.stringify(jwk),
  });
  const response = wasm.handleRequest(request);
  const parsed = JSON.parse(response);
  return parsed.signature;
};

export const verify = async (message: string, signature: string, jwk: any) => {
  const wasm = await init();
  const request = JSON.stringify({
    command: 'verify',
    message,
    signature,
    jwk: JSON.stringify(jwk),
  });
  const response = wasm.handleRequest(request);
  const parsed = JSON.parse(response);
  return parsed.verified;
};
