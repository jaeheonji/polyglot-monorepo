import { sayHello as sayHelloA } from '@polyglot-monorepo/packageA'
import { sayHello as sayHelloB } from '@polyglot-monorepo/packageB'

console.log(sayHelloA("I'm ServiceA"))
console.log(sayHelloB("I'm ServiceA"))
