import React, { useEffect, useState } from "react";
import { fetchUsers, createUser } from "../api";

const UsersList = () => {
    const [users, setUsers] = useState([]);

    useEffect(() => {
        fetchUsers()
            .then((response) => setUsers(response.data))
            .catch((error) => console.error("Error fetching users:", error));
    }, []);

    const handleCreateUser = () => {
        createUser()
            .then(() => fetchUsers().then((response) => setUsers(response.data)))
            .catch((error) => console.error("Error creating user:", error));
    };

    return (
        <div>
            <h1>Users</h1>
            <button onClick={handleCreateUser}>Create User</button>
            <ul>
                {users.map((user) => (
                    <li key={user.id}>
                        ID: {user.id}, Balance: {user.balance}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default UsersList;
