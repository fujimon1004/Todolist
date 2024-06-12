import React, { useState } from 'react';
import axios from 'axios';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';

interface Todo {
  id: string;
  task: string;
}

const fetchTodos = async () => {
  const { data } = await axios.get('http://localhost:8000/todos');
  return data;
};

const TodoApp: React.FC = () => {
  const queryClient = useQueryClient();
  const { data: todos, isLoading } = useQuery<Todo[]>({
    queryKey: ['todos'],
    queryFn: fetchTodos,
  });
  const mutation = useMutation({
    mutationFn: (newTodo: Todo) => axios.post('http://localhost:8000/todos', newTodo),
    onSuccess: () => {
      queryClient.invalidateQueries(['todos']);
    },
  });

  const [task, setTask] = useState('');

  const addTodo = () => {
    const newTodo = { id: Math.random().toString(36).substr(2, 9), task: task };
    mutation.mutate(newTodo);
    setTask('');
  };

  if (isLoading) return <div>Loading...</div>;

  return (
    <div>
      <h1>Todo List</h1>
      <input
        type="text"
        value={task}
        onChange={(e) => setTask(e.target.value)}
      />
      <button onClick={addTodo}>Add Todo</button>
      <ul>
        {todos?.map(todo => (
          <li key={todo.id}>{todo.task}</li>
        ))}
      </ul>
    </div>
  );
};

export default TodoApp;
