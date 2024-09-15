import React from 'react';

const Cart = () => {
  const items = [
    { id: 1, name: 'Штаны', price: 50 },
    { id: 2, name: 'Юбка', price: 40 },
  ];

  return (
    <div className="cart">
      <h1>Корзина</h1>
      <ul>
        {items.map(item => (
          <li key={item.id}>{item.name} - {item.price} $</li>
        ))}
      </ul>
    </div>
  );
};

export default Cart;