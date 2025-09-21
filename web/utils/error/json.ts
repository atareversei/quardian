export const ParsingJSONErrorName = 'ParsingJSONError';

export class JSONParsingError extends Error {
  constructor() {
    super("couldn't parse the JSON string.");
    this.name = ParsingJSONErrorName;
  }
}

export function isJSONParsingError(error: unknown): error is JSONParsingError {
  return error instanceof JSONParsingError && error.name === ParsingJSONErrorName;
}
