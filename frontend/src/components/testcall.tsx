import React, { useState } from "react";
import axios from "axios";

//COMPONENT TO TEST END TO END CALLS to DB

// Define the TypeScript interface for a listing
interface Listing {
  email: string;
  listID: number;
  title: string;
  description: string;
  img: string;
  active: boolean;
  category: string;
  rarity: string;
}

const Listings: React.FC = () => {
  // State to store listings fetched from the server
  const [listings, setListings] = useState<Listing[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(false);

  // Function to fetch listings from the backend using axios
  const fetchListings = async () => {
    setLoading(true); // Set loading to true when request is made
    setError(null); // Clear previous errors
    try {
      const response = await axios.get("http://localhost:8080/allList");
      setListings(response.data);
    } catch (err) {
      setError("Failed to fetch listings REACOM");
    } finally {
      setLoading(false); // Set loading to false after request is complete
    }
  };

  return (
    <div>
      <h1>Listings</h1>
      <button onClick={fetchListings} disabled={loading}>
        {loading ? "Loading..." : "Fetch Listings"}
      </button>

      {error && <p style={{ color: "red" }}>{error}</p>}

      {listings.length === 0 && !loading && !error ? (
        <p>No listings available. Click the button to fetch.</p>
      ) : (
        <ul>
          {listings.map((listing) => (
            <li key={`${listing.email}-${listing.listID}`}>
              <h2>{listing.title}</h2>
              <p>
                <strong>Description:</strong> {listing.description}
              </p>
              <p>
                <strong>Email:</strong> {listing.email}
              </p>
              <p>
                <strong>Category:</strong> {listing.category}
              </p>
              <p>
                <strong>Rarity:</strong> {listing.rarity}
              </p>
              <p>
                <strong>Active:</strong> {listing.active ? "Yes" : "No"}
              </p>
              <img
                src={listing.img}
                alt={listing.title}
                style={{ width: "200px" }}
              />
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default Listings;
