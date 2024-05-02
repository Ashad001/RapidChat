import React, { useState } from 'react';
import EmojiPicker from 'emoji-picker-react'; // Ensure this is correctly imported

function SelectEmoji() {
  const [selectedEmoji, setSelectedEmoji] = useState(null);

  function handleEmojiSelect(emoji) {
    setSelectedEmoji(emoji);
  }

  return (
    <div>
      <EmojiPicker
      onEmojiSelect={handleEmojiSelect}
      emojiSize={4}
      height={300}
      width={250}
      theme='dark'
      emojiStyle='facebook'
      reactionsDefaultOpen={true}
      allowExpandReactions={false}
      />
      <h1>Selected Emoji: {selectedEmoji}</h1>
    </div>
  );
}

export default SelectEmoji;
