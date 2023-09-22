import { Sequelize } from "sequelize";

let {
  DB_NAME: db_name,
  DB_USER: db_user,
  DB_PASSWORD: db_password,
} = process.env;

console.log(process.env.DB_USER);

console.log(db_name, db_user, db_password);

const sequelize = new Sequelize(db_name!, db_user!, db_password!, {
  host: "localhost",
  dialect: "postgres",
});

export default sequelize;
