import { useCallback, useEffect, useState } from "react";
import { Gift } from "../types";
import api from "../api";

export const useGift = (id?: string) => {
  const [gift, setGift] = useState<Gift>();
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<boolean>(false);

  const loadGift = useCallback(async () => {
    if (!id) return;

    setLoading(true);

    try {
      const gift = await api.getGift(id);
      setGift(gift);
    } catch (e) {
      setError(true);
    } finally {
      setLoading(false);
    }
  }, [id]);

  useEffect(() => {
    loadGift();
  }, [loadGift, id]);

  return { gift, loadGift, loading, error };
}
