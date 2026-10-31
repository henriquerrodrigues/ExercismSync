export function decodedValue(color: string[]): number {
  const first = COLORS.indexOf(color[0]);
  const second = COLORS.indexOf(color[1]);
  const textNumber = `${first}${second}` 
  return Number(textNumber)
}

export const COLORS: string[] = [
  "black",
  "brown",
  "red",
  "orange",
  "yellow",
  "green",
  "blue",
  "violet",
  "grey",
  "white",
];
