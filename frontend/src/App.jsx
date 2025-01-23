import { useEffect } from "react";
import Home from "./components/pages/home";
import axios from "axios";

const App = () => {

  // TESTANDO API GO
  // REQUEST GET
  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get("http://localhost:8000", {
          withCredentials: true, // Envia o cookie de sessão
          headers: {
            "Content-Type": "application/json",
          },
        });
        console.log(response.data);
      } catch (error) {
        console.error("Erro ao buscar dados:", error);
      }
    };

    fetchData();
  }, []);

  return (
    <>
      <Home />
    </>
  );
};

export default App;