import { createClient } from 'redis';

const client = createClient({
    url: 'redis://localhost:6379'
});

await client.connect();

// Strings
await client.set('name', 'Miroslav');
const name = await client.get('name');
console.log(name);  // Miroslav

// Hash
await client.hSet('user:1', { name: 'Alice', age: '30' });
const user = await client.hGetAll('user:1');
console.log(user);  // { name: 'Alice', age: '30' }

// Lists
await client.lPush('mylist', 'a', 'b', 'c');
const list = await client.lRange('mylist', 0, -1);
console.log(list);  // [ 'c', 'b', 'a' ]

await client.quit();
