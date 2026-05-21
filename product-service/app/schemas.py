from pydantic import BaseModel


class ProductCreate(BaseModel):
    name: str
    description: str
    price: float
    stock: int


class ProductResponse(ProductCreate):
    id: int
    status: str

    class Config:
        from_attributes = True