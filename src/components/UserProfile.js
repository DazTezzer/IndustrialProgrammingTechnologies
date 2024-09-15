import React from 'react';

const UserProfile = () => {
  return (
    <div className="user-profile">
      <h1>Личный кабинет</h1>
      <form>
        <input type="text" placeholder="Имя" />
        <input type="email" placeholder="Электронная почта" />
        <button type="submit">Сохранить</button>
      </form>
    </div>
  );
};

export default UserProfile;