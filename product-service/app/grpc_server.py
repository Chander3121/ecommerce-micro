import grpc

from concurrent import futures

from app.database import SessionLocal

from app import models

from app import product_pb2
from app import product_pb2_grpc


class ProductService(
    product_pb2_grpc.ProductServiceServicer
):

    def GetProduct(self, request, context):
        db = SessionLocal()

        product = db.query(models.Product).filter(
            models.Product.id == request.id
        ).first()

        if not product:
            context.set_code(
                grpc.StatusCode.NOT_FOUND
            )

            context.set_details(
                "Product not found"
            )

            return product_pb2.ProductResponse()

        return product_pb2.ProductResponse(
            id=product.id,
            name=product.name,
            description=product.description,
            price=product.price,
            stock=product.stock,
            status=product.status
        )


def serve():
    print("Starting gRPC server...")
    server = grpc.server(
        futures.ThreadPoolExecutor(max_workers=10)
    )

    product_pb2_grpc.add_ProductServiceServicer_to_server(
        ProductService(),
        server
    )

    server.add_insecure_port("[::]:50051")

    server.start()

    print("gRPC Product Server running on 50051")

    server.wait_for_termination()
