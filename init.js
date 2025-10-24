

db.getSiblingDB('admin').createUser({
  user: 'admin',
  pwd: '868326481923',
  roles: [
    { role: 'root', db: 'admin' } 
  ]
});



db.getSiblingDB('myapp');


db.getSiblingDB('myapp').createUser({
  user: 'app_user', 
  pwd: 'passwordQ123',
  roles: [
    { role: 'readWrite', db: 'myapp' },    
    { role: 'dbAdmin', db: 'myapp' }       
  ]
});

db = db.getSiblingDB('myapp');
db.createCollection('users');
db.createCollection('products');


