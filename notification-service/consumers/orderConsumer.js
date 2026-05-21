const amqp = require("amqplib");

async function consumeOrders() {
  try {
    const connection = await amqp.connect(
      process.env.RABBITMQ_URL
    );

    const channel = await connection.createChannel();

    const queue = "order.created";

    await channel.assertQueue(queue, {
      durable: true,
    });

    console.log(
      "Waiting for order events..."
    );

    channel.consume(queue, (message) => {
      if (message !== null) {
        const order = JSON.parse(
          message.content.toString()
        );

        console.log(
          "Order event received:"
        );

        console.log(order);

        // Future:
        // send email
        // send SMS
        // push notification

        channel.ack(message);
      }
    });
  } catch (error) {
    console.error(error);
  }
}

module.exports = consumeOrders;
