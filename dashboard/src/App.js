import './App.css';
import MainWrapper from './components/MainWrapper';
import Dashboard from './components/dashboard';
import Navbar from './components/navbar';

function App() {
    return (
        <MainWrapper>
            <Navbar />
            <Dashboard />
        </MainWrapper>
    );
}

export default App;
