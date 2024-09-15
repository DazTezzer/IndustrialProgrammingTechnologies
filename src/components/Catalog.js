import React from 'react';
import ProductCard from './ProductCard.js';

const Catalog = () => {
  const products = [
    { id: 1, name: 'Штаны', category: 'Мужская', price: 50 },
    { id: 2, name: 'Юбка', category: 'Женская', price: 40 },
    { id: 3, name: 'Рубашка', category: 'Женская', price: 30 },
    // Добавьте больше товаров по необходимости
  ];

  return (
    <div className="catalog">
      <h1>Каталог товаров</h1>
      <div className="product-list">
        {products.map(product => (
          <ProductCard key={product.id} product={product} />
        ))}
      </div>
    </div>
  );
};

export default Catalog;