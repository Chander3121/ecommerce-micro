require("dotenv").config();

const express = require("express");

const consumeOrders = require(
  "./consumers/orderConsumer"
);

const app = express();

app.get("/", (req, res) => {
  res.json({
    status:
      "Notification Service Running",
  });
});

consumeOrders();

app.listen(3004, () => {
  console.log(
    "Notification Service running on 3004"
  );
});
