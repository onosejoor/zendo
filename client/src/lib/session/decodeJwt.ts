export function decodeJwt(token: string) {
  const [header, payload] = token.split(".");
  const decodedHeader = JSON.parse(atob(header));
  const decodedPayload = JSON.parse(atob(payload));

  
  return { header: decodedHeader, payload: decodedPayload };
}
