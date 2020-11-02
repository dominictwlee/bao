import { sum } from './sum';

export function greet(name) {
  dummy();
  return `Hello ${name}`;
}

export function dummy() {
  console.log(greet('dummy'));
  console.log(sum(2, 5));
}