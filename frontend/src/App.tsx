import React from 'react';
import TodoApp from './Todo';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const queryClient = new QueryClient();

const App: React.FC = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <div className="App">
        <TodoApp />
      </div>
    </QueryClientProvider>
  );
};

export default App;
