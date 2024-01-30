import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";

function App() {
	const [count, setCount] = useState(0);

	const increment = () => {
		setCount(count + 1);
	};

	const decrement = async () => {
		setCount(count - 1);
	};

	const saveCount = async () => {
		try{
			const ret = await fetch("/api/count", {
				method: "POST",
				body: JSON.stringify({count})
			})
			const data = await ret.json()
			if(!data?.success){
				console.error(data.error)
				return
			}
			
			setCount(data.newCount)
		}catch(e){
			console.error(e)
		}
	}

	const loadCount = async () => {
		const ret = await fetch("/api/count", {
			method: "GET"
		});
		const data = await ret.json()
		if(!data.success){
			console.error(data.error)
			return
		}
		console.log(data)
		setCount(data.count)
	}

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
			<div className="card">
				<div className="horizon">
					<button type="button" onClick={decrement}>
						-
					</button>
					<strong>
						COUNT IS " <u>{count}</u> "
					</strong>
					<button type="button" onClick={increment}>
						+
					</button>
				</div>
			</div>
			<div className="horizon">
			<button type="button" onClick={loadCount}>Load from DB</button>
				<button type="button" onClick={saveCount}>Save to DB</button>
			</div>
			<p className="read-the-docs">
				Click on the Vite and React logos to learn more
			</p>
		</>
	);
}

export default App;
