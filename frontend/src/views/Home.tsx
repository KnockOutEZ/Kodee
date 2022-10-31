import React, { useEffect } from 'react'
import { Link } from 'react-router-dom';
// import icon from '../../../assets/icon.svg';
import '../../wailsjs/go/main/App'
import { Greet, Notification,GetCpuUsage,GetRamUsage,GetBandwithSpeed } from '../../wailsjs/go/main/App';
import '../styles/home.css'

const Home = () => {
    useEffect(() => {
      (async () => {
        console.log(await Notification("nigga","chigga"))
        console.log(await Greet("nigga"))
    })();
      // eslint-disable-next-line promise/catch-or-return
      // run function every 5 seconds
      // setInterval(getNotification,5000)
    })

    async function getUsage(){
      console.log(await GetCpuUsage())
      console.log(await GetRamUsage())
      console.log(await GetBandwithSpeed())
    }

    return (
        <div className='body'>
            <div className="Hello">
            {/* <img width="200" alt="icon" src={icon} /> */}
          </div>
          <h1>KODEE</h1>
          <div className="Hello">
            <a
              href="https://electron-react-boilerplate.js.org/"
              target="_blank"
              rel="noreferrer"
            >
              <button className='button' type="button">
                <span role="img" aria-label="books">
                  📚
                </span>
                Read our docs
              </button>
            </a>
              <button className='button' type="button" onClick={getUsage}>
                <span role="img" aria-label="folded hands">
                  🙏
                </span>
                Donate
              </button>
              
              <Link to="/todo">
                <button className='button' type="button">
                  TODO
                </button>
              </Link>
          </div>
        </div>
      )}

export default Home