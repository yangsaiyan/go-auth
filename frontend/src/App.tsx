import { BrowserRouter, Routes, Route } from "react-router";
import "./App.css";
import Home from "./pages/home/Home";
import Auth from "./pages/auth/Auth";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/auth" element={<Auth />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
