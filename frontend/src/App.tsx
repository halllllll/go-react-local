import { useState } from "react";
import { usePostCount } from "./service/count";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import { CountListComponent } from "./counttable";

function App() {
  const [count, setCount] = useState(0);
  const { mutate: postCountMutate } = usePostCount();
  const [showList, setShowList] = useState<boolean>(false);
  const [showErrorMsg, setShowErrMsg] = useState<boolean>(false);

  const increment = (val: number | any) => {
    if (typeof val === "number") {
      setCount(count + val);
    } else {
      setCount(count + 1);
    }
  };

  const decrement = (val: number | any) => {
    if (typeof val === "number") {
      setCount(count - val);
    } else {
      setCount(count - 1);
    }
  };

  const saveCount = async () => {
    postCountMutate(
      { count },
      {
        onSettled: () => {
          console.log("fire");
        },
        onSuccess: () => {
          console.log("done!");
          setShowErrMsg(false);
        },
        onError: (err) => {
          console.error("failed...");
          console.log(err);
          setShowErrMsg(true);
        },
      }
    );
  };

  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank" rel="noreferrer">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank" rel="noreferrer">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <h2>{">> on Go(Gin) local server <<"}</h2>
      <div className="card">
        <div className="horizon">
          <button type="button" onClick={() => decrement(10)}>
            {"-10"}
          </button>
          <button type="button" onClick={decrement}>
            -
          </button>
          <strong>
            " <u>{count}</u> "
          </strong>
          <button type="button" onClick={increment}>
            +
          </button>
          <button type="button" onClick={() => increment(10)}>
            {"10"}
          </button>
        </div>
      </div>
      <div className="card">
        {showErrorMsg && (
          <div className="error-message">サーバーエラーです</div>
        )}
        <button type="button" onClick={saveCount}>
          Save to DB
        </button>
      </div>
      <div className="card">
        <button type="button" onClick={() => setShowList(!showList)}>
          {showList ? "close" : "Show List"}
        </button>
      </div>
      {showList && <CountListComponent />}
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}

export default App;
