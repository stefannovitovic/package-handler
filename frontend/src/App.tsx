import { FormEvent, useState } from "react";
import "./App.css";

interface PackResult {
  total_items: number;
  packs: Record<string, number>;
}

export default function App() {
  const [order, setOrder] = useState("4200");
  const [packSizes, setPackSizes] = useState("250, 500, 1000, 2000, 5000");
  const [result, setResult] = useState<PackResult | null>(null);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  function parseSizes(text: string): number[] {
    return text
      .split(/[\s,]+/)
      .filter(Boolean)
      .map((s) => {
        const n = Number(s);
        if (!Number.isInteger(n) || n <= 0)
          throw new Error(`Invalid pack size: "${s}"`);
        return n;
      });
  }

  async function handleSubmit(e: FormEvent) {
    e.preventDefault();
    setResult(null);
    setError("");

    const orderNum = Number(order);
    if (!Number.isInteger(orderNum) || orderNum <= 0) {
      setError("Order must be a positive integer.");
      return;
    }

    let sizes: number[];
    try {
      sizes = parseSizes(packSizes);
      if (sizes.length === 0) throw new Error("Enter at least one pack size.");
    } catch (err: any) {
      setError(err.message);
      return;
    }

    setLoading(true);
    try {
        const base = import.meta.env.VITE_BACKEND_URL ?? "";
        const res = await fetch(`${base}/calculate-packs`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ order: orderNum, pack_sizes: sizes }),
      });
      if (!res.ok) {
        const text = await res.text();
        setError(`${res.status} — ${text.trim() || res.statusText}`);
        return;
      }
      setResult(await res.json());
    } catch (err: any) {
      setError("Request failed: " + err.message);
    } finally {
      setLoading(false);
    }
  }

  const totalPacks = result
    ? Object.values(result.packs).reduce((a, b) => a + b, 0)
    : 0;

  return (
    <div className="shell">
      <header>
        <h1>Package Handler</h1>
        <p className="subtitle">
          Calculate the optimal pack combination for an order.
        </p>
      </header>

      <form onSubmit={handleSubmit}>
        <label>
          <span>Order (items)</span>
          <input
            type="number"
            min={1}
            value={order}
            onChange={(e) => setOrder(e.target.value)}
            required
          />
        </label>

        <label>
          <span>Pack sizes (comma-separated)</span>
          <textarea
            rows={2}
            value={packSizes}
            onChange={(e) => setPackSizes(e.target.value)}
            required
          />
        </label>

        <button type="submit" disabled={loading}>
          {loading ? "Calculating…" : "Calculate"}
        </button>
      </form>

      {error && <div className="card error">{error}</div>}

      {result && (
        <div className="card result">
          <div className="stat-row">
            <div className="stat">
              <span className="stat-value">{result.total_items}</span>
              <span className="stat-label">Items shipped</span>
            </div>
            <div className="stat">
              <span className="stat-value">{totalPacks}</span>
              <span className="stat-label">Packs used</span>
            </div>
            <div className="stat">
              <span className="stat-value">
                +{result.total_items - Number(order)}
              </span>
              <span className="stat-label">Over-ship</span>
            </div>
          </div>

          <h3>Breakdown</h3>
          <table>
            <thead>
              <tr>
                <th>Pack size</th>
                <th>Qty</th>
                <th>Subtotal</th>
              </tr>
            </thead>
            <tbody>
              {Object.entries(result.packs)
                .sort(([a], [b]) => Number(b) - Number(a))
                .map(([size, qty]) => (
                  <tr key={size}>
                    <td>{Number(size).toLocaleString()}</td>
                    <td>×{qty}</td>
                    <td>{(Number(size) * qty).toLocaleString()}</td>
                  </tr>
                ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
