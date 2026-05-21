const express = require("express");

const cors = require("cors");

const morgan = require("morgan");

const {
  createProxyMiddleware,
} = require("http-proxy-middleware");

const app = express();

app.use(cors());

app.use(morgan("dev"));


// AUTH SERVICE

app.use(
  "/auth",
  createProxyMiddleware({
    target: "http://auth-service:3001",

    changeOrigin: true,

    pathRewrite: {
      "^/auth": "",
    },
  })
);


// PRODUCT SERVICE

app.use(
  "/api/products",
  createProxyMiddleware({
    target: "http://product-service:3002/products",

    changeOrigin: true,
  })
);


// ORDER SERVICE

app.use(
  "/api/orders",
  createProxyMiddleware({
    target: "http://order-service:3003/orders",

    changeOrigin: true,
  })
);

app.get("/", (req, res) => {
  res.json({
    status: "API Gateway Running",
  });
});


app.listen(8080, () => {
  console.log(
    "API Gateway running on 8080"
  );
});
