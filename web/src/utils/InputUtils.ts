export const validLettersAndNumbers = (
  input: string,
  allowSpaces: boolean = false
): boolean => {
  if (!input) return false;
  const regex = allowSpaces ? /^[a-zA-Z0-9\s]*$/ : /^[a-zA-Z0-9]*$/;
  return regex.test(input);
};

export const validLetters = (
  input: string,
  allowSpaces: boolean = false
): boolean => {
  if (!input) return false;
  const regex = allowSpaces ? /^[a-zA-ZÀ-ÖØ-öø-ÿ\s]*$/ : /^[a-zA-ZÀ-ÖØ-öø-ÿ]*$/;
  return regex.test(input);
};

export const validNumbers = (input: string): boolean => {
  if (!input) return false;
  const regex = /^[0-9]*$/;
  return regex.test(input);
};

export const validEmail = (email: string): boolean => {
  if (!email) return false;
  const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return regex.test(email);
};

export const validWithNoSpaces = (input: string): boolean => {
  if (!input) return false;
  const regex = /\s/;
  return !regex.test(input);
};
