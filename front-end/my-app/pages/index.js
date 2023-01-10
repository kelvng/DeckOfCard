import { useState, useEffect } from 'react';
import Link from 'next/link';
import axios from 'axios';
import styles from '../styles/Home.module.css';
import {error} from "next/dist/build/output/log";

async function getPlayer(playerID) {
  try {
    const response = await axios.get(`https://test-golang.herokuapp.com/player/${playerID}`);
    return response.data;
  } catch (error) {
    console.error(error);
  }
}

async function shuffleDeck() {
  try {
    const response = await axios.post(`https://test-golang.herokuapp.com/shuffle`);
    setDeck(response.data);
  } catch (error) {
    console.error(error);
  }
}

async function drawCards() {
  try {
    const response = await axios.delete(`https://test-golang.herokuapp.com/drawn`);
    setGameResult(response.data);
  } catch (error) {
    console.error(error);
  }
}

async function resetGame() {
  try {
    await axios.post(`https://test-golang.herokuapp.com/new`);
    setDeck([]);
    setGameResult(null);
  } catch (error) {
    console.error(error);
  }
}

const valueMap = {
  'Two': 2,
  'Three': 3,
  'Four': 4,
  'Five': 5,
  'Six': 6,
  'Seven': 7,
  'Eight': 8,
  'Nine': 9,
  'Ten': "J",
  'Jack': "J",
  'Queen': "Q",
  'King': "K",
  'Ace': "A",
}

function convertCardToString(card) {
  const value = valueMap[card.value] || card.value;
  const suit = card.suit[0];
  return `${value}${suit}`;
}

export default function Home() {
  const [player, setPlayer] = useState(null);
  const [deck, setDeck] = useState([]);
  const [gameResult, setGameResult] = useState(null);
  const [playerID, setPlayerID] = useState(1);

  useEffect(() => {
    async function fetchData() {
      const player = await getPlayer(playerID);
      setPlayer(player);
    }
    fetchData();
  }, [playerID]);

  useEffect(() => {
    async function fetchData() {
      const response = await axios.get(`https://test-golang.herokuapp.com/top`);
      setDeck(response.data);
    }
    fetchData();
  }, [deck]);

  return (

      <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }} className={styles.main}>
        {player ? (
            <div>
              <p>ID: {player.id}</p>
              <p>Name: {player.name}</p>
              <p>Coins: {player.coins}</p>
              <div >
                <p>Hand:</p>
                <div  style={{ display: 'flex', flexWrap: 'wrap' }}>
                  {player && player.hand ? player.hand.map(card => {
                    const cardString = convertCardToString({ value: card.value, suit: card.suit });
                    return (
                        <div key={cardString}>
                          <img src={`https://deckofcardsapi.com/static/img/${cardString}.png`} alt={`${card.value} of ${card.suit}`} style={{ width: '100px', height: '150px' }}/>
                        </div>
                    )
                  }) : null}
                </div>

              </div>
              <p>Points: {player.points}</p>



            </div>
        ) : (
            <p>Loading...</p>
        )}
        <div>
        </div>

        {gameResult ? <p>{gameResult}</p> : null}

        <div id="status-box">
          Status: {error}
        </div>


        <div>
          <Link href="/">
            <button className={styles.prevButton} onClick={() => setPlayerID(playerID - 1)}>Prev</button>
          </Link>
          <Link href="/">
            <button className={styles.nextButton} onClick={() => setPlayerID(playerID + 1)}>Next</button>
          </Link>
        </div>

        <div>
          <button className={styles.button} onClick={shuffleDeck} >Shuffle</button>
          <button className={styles.button} onClick={drawCards}>Draw</button>
          <button className={styles.button} onClick={() => {
            setGameResult(gameResult);
            setPlayerID(playerID);
          }}>Reveal</button>
          <button className={styles.button} onClick={resetGame}>Reset</button>
        </div>

      </div>
  )};