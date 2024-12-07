import axios from "axios";

const API_URL = "http://localhost:8081";

// Получение всех пользователей
export const fetchUsers = () => axios.get(`${API_URL}/users`);

// Создание нового пользователя
export const createUser = () => axios.post(`${API_URL}/users`);

// Получение баланса пользователя
export const getUserBalance = (userId) =>
    axios.get(`${API_URL}/users/${userId}`);

// Создание транзакции
export const createTransaction = (data) =>
    axios.post(`${API_URL}/transactions`, data);

// Получение транзакций пользователя
export const getUserTransactions = (userId) =>
    axios.get(`${API_URL}/transactions/${userId}`);
