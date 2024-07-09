// App.js 中可以不显式的导入 React
import { Routes, Route } from "react-router-dom";

import Home from "./page/Home";
import Blog from "./page/Blog";

import 'bootstrap/dist/css/bootstrap.min.css';
import "./App.css";
// 在 React 项目中引入 Bootstrap 的 CSS 样式文件，让 Bootstrap 的所有预定义样式生效。
// 在 React 组件中直接使用 Bootstrap 提供的样式和组件。

function App() {
  

  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/blog/:id" element={<Blog />} />
    </Routes>

  );
}

export default App;
