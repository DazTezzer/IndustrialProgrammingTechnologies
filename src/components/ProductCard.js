import React from 'react';

const ProductCard = ({ product }) => {
  const handleAddToCart = () => {
    // Логика добавления в корзину
    alert(`Товар ${product.name} добавлен в корзину!`);
  };

  return (
    <div className="product-card">
      <h2>{product.name}</h2>
      <p>Цена: {product.price} $</p>
      <button onClick={handleAddToCart}>Добавить в корзину</button>
    </div>
  );
};

export default ProductCard;