import { JSONParsingError } from '../error/json';

/**
 * Parses the `Response` object as a JSON string and returns an object.
 *
 * @param res - `Response` object.
 * @returns JSON.
 * @throws {ParsingJSONError} If the received JSON is invalid.
 */
export async function parseJSONResponse(res: Response): Promise<{ [key: string]: any }> {
  let json = null;
  try {
    json = await res.json();
    return json;
  } catch (err) {
    throw new JSONParsingError();
  }
}
