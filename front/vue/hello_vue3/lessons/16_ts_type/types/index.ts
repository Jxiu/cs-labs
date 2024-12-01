// ts 的接口用于限制对象的具体属性，有点像java类
export interface Person {
    id: string,
    name: string,
    age: number
}
// 自定义类型
export type PersonList = Array<Person>
export type PersonArray = Person[]