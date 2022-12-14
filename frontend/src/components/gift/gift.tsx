import React, {useState} from "react";
import {Params, useParams} from "react-router-dom";
import emoji from 'emoji-dictionary';
import {useGift} from "../../hooks/useGift";
import Loader from "../loader/loader";
import api from "../../api";
import './gift.scss';

export const DisplayGift = () => {
  const {id} = useParams<Params>();
  const {gift, loadGift, loading, error} = useGift(id);

  const onClickUse = async () => {
    if (!gift) return;

    await api.postUseGift(gift);

    loadGift();
  }

  if (error) {
    return (
      <div>Gift not found.</div>
    )
  }

  if (loading || !gift) {
    return (
      <Loader />
    );
  }

  const getEmoji = (name: string) => (
    emoji.getUnicode(name)
  );

  const Button = ({i}: { i: number }) => {
    const [pressed, setPressed] = useState<boolean>(false);

    const onClick = () => {
      setPressed(true);
      setTimeout(() => onClickUse(), 350)
    }

    const className = `gift-button${pressed ? ' shrinking' : ''}`;

    return (
      <div key={i} className={className} role="button" tabIndex={0} onClick={onClick}>
        {getEmoji(gift.displayText)}
      </div>
    )
  }

  return (
    <div className="container">
      <div className="gift">
        <h1>{gift.name}</h1>
        <p>{gift.amount > 0 ? gift.description: `Ingen flere ${getEmoji(gift.displayText)} igjen :(`}</p>
        <div className="gift-buttons">
          {[...Array(gift.amount)].map((_, i) => <Button i={i}/>)}
        </div>
      </div>
    </div>
  )
}