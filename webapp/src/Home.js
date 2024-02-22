import React from 'react';
import { Button, Container, Card } from 'react-bootstrap'; // Example using Bootstrap

function Home() {
  return (
    <Container>
      <Card.Header>Welcome to My Website!</Card.Header>
      <p>This is a simple home page using a component library.</p>
      <Button variant="primary">Learn More</Button>
    </Container>
  );
}

export default Home;