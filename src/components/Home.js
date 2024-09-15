import React from 'react';

const Home = () => {
  return (
    <div className="home">
      <h1>Добро пожаловать в интернет-магазин одежды!</h1>
      <div className="categories">
        <h2>Выбор категории:</h2>
        <button>Женская одежда</button>
        <button>Мужская одежда</button>
        <button>Детская одежда</button>
        <button>Ботинки</button>
        <button>Шляпы</button>
      </div>
    </div>
  );
};

export default Home;