function addQuotes(value: any) {
  if (typeof value === 'string') return `'${value}'`;
  return value;
}

export function printObject(obj?: Record<string, any>): string {
  if (!obj) return 'undefined';
  let objstr = '{';
  Object.entries(obj).forEach(([key, value], i, arr) => {
    objstr += `'${key}': ${addQuotes(value)}${arr.length - 1 === i ? '' : ','}`;
  });
  objstr += '}';
  return objstr;
}
