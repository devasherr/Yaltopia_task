import { Routes, Route } from "react-router-dom";
import Dashboard from "./pages/Dashboard";
import CricketFixture from "./pages/CricketFixture";
import VolleyFixture from "./pages/VolleyFixture";

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Dashboard />} />
      <Route path="/fixture/cricket/:id" element={<CricketFixture />} />
      <Route path="/fixture/volleyball/:id" element={<VolleyFixture />} />
    </Routes>
  );
}
