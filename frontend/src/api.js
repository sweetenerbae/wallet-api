import axios from "axios";

const API_URL = "http://localhost:8081";

// api.js

export const fetchUsers = async () => {
    const response = await fetch("http://localhost:8081/users");
    if (!response.ok) {
        throw new Error("Error fetching users");
    }
    return await response.json();
};

export const createUser = async (user) => {
    const response = await fetch("http://localhost:8081/users", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(user),
    });
    if (!response.ok) {
        throw new Error("Error creating user");
    }
    return await response.json();
};

// Получение баланса пользователя
export const getUserBalance = (userId) =>
    axios.get(`${API_URL}/users/${userId}`);

// Создание транзакции
export const createTransaction = (data) =>
    axios.post(`${API_URL}/transactions`, data);

// Получение транзакций пользователя
export const getUserTransactions = (userId) =>
    axios.get(`${API_URL}/transactions/${userId}`);
