from fastapi import FastAPI, Depends, HTTPException
import threading

from app.database import engine, get_db, Base
from app import models, schemas, crud
from app.grpc_server import serve

from sqlalchemy.orm import Session

Base.metadata.create_all(bind=engine)

app = FastAPI()


@app.on_event("startup")
async def startup_event():
    print("STARTING GRPC THREAD")

    grpc_thread = threading.Thread(
        target=serve,
        daemon=True
    )

    grpc_thread.start()


@app.get("/")
def health():
    return {
        "status": "Product Service Running"
    }


@app.post("/products", response_model=schemas.ProductResponse)
def create_product(
    product: schemas.ProductCreate,
    db: Session = Depends(get_db)
):
    return crud.create_product(db, product)


@app.get("/products")
def products(db: Session = Depends(get_db)):
    return crud.get_products(db)


@app.get("/products/{product_id}")
def product(product_id: int, db: Session = Depends(get_db)):
    product = crud.get_product(db, product_id)

    if not product:
        raise HTTPException(
            status_code=404,
            detail="Product not found"
        )

    return product
