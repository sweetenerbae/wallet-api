import React, { useEffect, useState } from "react";
import { fetchUsers, createUser } from "../api";

const UsersList = () => {
    const [users, setUsers] = useState([]);
    const [name, setName] = useState("");  // состояние для имени пользователя
    const [balance, setBalance] = useState(0);  // состояние для баланса

    useEffect(() => {
        fetchUsers()
            .then((response) => setUsers(response.data))
            .catch((error) => console.error("Error fetching users:", error));
    }, []);

    const handleCreateUser = () => {
        if (name && balance >= 0) {  // проверка на пустое имя и отрицательный баланс
            createUser({ name, balance })  // передаем имя и баланс в функцию createUser
                .then(() => {
                    fetchUsers().then((response) => setUsers(response.data));  // обновляем список пользователей
                    setName("");  // очищаем поля формы после успешного создания
                    setBalance(0);
                })
                .catch((error) => console.error("Error creating user:", error));
        } else {
            alert("Please provide a valid name and balance.");
        }
    };

    return (
        <div>
            <h1>Users</h1>

            {/* Форма для ввода данных нового пользователя */}
            <div>
                <input
                    type="text"
                    placeholder="Name"
                    value={name}
                    onChange={(e) => setName(e.target.value)}  // обновляем имя
                />
                <input
                    type="number"
                    placeholder="Balance"
                    value={balance}
                    onChange={(e) => setBalance(e.target.value)}  // обновляем баланс
                />
                <button onClick={handleCreateUser}>Create User</button>
            </div>

            <ul>
                {users.map((user) => (
                    <li key={user.id}>
                        ID: {user.id}, Name: {user.name}, Balance: {user.balance}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default UsersList;
