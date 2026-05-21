from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy.orm import Session

from app import models, schemas, crud

from app.database import engine, get_db, Base

Base.metadata.create_all(bind=engine)

app = FastAPI()


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
