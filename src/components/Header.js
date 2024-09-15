import React from 'react';
import { Link } from 'react-router-dom';
import './Header.css';

const Header = () => {
  return (
    <header className="header">
      <nav>
        <Link to="/">Главная</Link>
        <Link to="/catalog">Каталог</Link>
        <Link to="/profile">Личный кабинет</Link>
        <Link to="/cart">Корзина</Link>
      </nav>
    </header>
  );
};

export default Header;