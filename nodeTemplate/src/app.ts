import express from "express";
import UserRoutes from "@routes/user.ts";
import sequelize from "./models/index.ts";

const app = express();

const port = 8080; // default port to listen

app.use("/user", UserRoutes);

// booting up sequelize
async () => {
  try {
    await sequelize.authenticate();
    console.log("Connection has been established successfully.");
  } catch (error) {
    console.error("Unable to connect to the database:", error);
  }
};

// start the Express server
app.listen(port, () => {
  console.log(`server started at http://localhost:${port}`);
});
