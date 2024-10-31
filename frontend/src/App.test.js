import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import App from './App';

global.fetch = jest.fn(() =>
  Promise.resolve({
    json: () => Promise.resolve({ message: 'Hello from Go!' }),
  })
);

test('renders message from API', async () => {
  render(<App />);
  const messageElement = await waitFor(() => screen.getByText(/Hello from Go!/i));
  expect(messageElement).toBeInTheDocument();
});
