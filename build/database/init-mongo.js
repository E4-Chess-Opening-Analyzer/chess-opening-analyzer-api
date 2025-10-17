db = db.getSiblingDB('goapi');

db.createUser({
  user: "ugoapi",
  pwd: "pgoapi",
  roles: [
    { role: "readWrite", db: "goapi" }
  ]
});
