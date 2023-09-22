import "module-alias/register";
import express from "express";
import UserRoutes from "@routes/user";

const app = express();

const port = 8080; // default port to listen

app.use("/user", UserRoutes);

// start the Express server
app.listen(port, () => {
  console.log(`server started at http://localhost:${port}`);
});
