package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Milua25/go_social/internal/store"
)

var usernames = []string{
	// Tech & Minimalist
	"pixel_perfect", "code_mist", "snack.exe", "html.cutie", "void_packet",
	"glitch_muse", "binary_bloom", "null_pointer", "cyber_niche", "bit_spirit",

	// Aesthetic & Nature
	"velvet_mood", "cloud_berry", "petal_cloud", "ivy_whisper", "moon_kissed",
	"lilac_haze", "moss_let", "solar_pulse", "star_drizzle", "wild_witchlet",

	// Creative & Artistic
	"ink_writer", "palette_dreams", "canvas_soul", "mistiq_flow", "auralla_art",
	"dusk_pulse", "neon_nymph", "abstract_mind", "color_story", "lyric_lune",

	// Trendy & Gen-Z (2026 Vibes)
	"delulu_girl", "zero_joy_club", "sad_scroller", "iced_matcha", "cringe_energy",
	"lowkey_legend", "not_feral", "moody_noodles", "offline_fairy", "serotonin_rush",

	// Gamers & Action
	"aim_ghost", "frag_spark", "respawn_dream", "next_rapid", "shadow_faerie",
	"stunlocker", "chaos_lord", "vortex_rider", "stealth_script", "recoil_roam",
}

var blogTitles = []string{
	"How I Rebuilt My Workflow with AI",     // Personal transformation
	"10 Habits of Digital Minimalists",      // Wellness trend
	"The Ethics of AI No One Talks About",   // Deep-dive topic
	"How to Scale Without Burning Out",      // Professional growth
	"Why Your 2026 Budget Isn't Working",    // Finance solution
	"The Truth About Passive AI Income",     // Money-making
	"Digital Detoxing for Remote Workers",   // Lifestyle shift
	"5 Mistakes That Kill Small Businesses", // Problem/Solution
	"Eco-Tech You Can Actually Use Today",   // Sustainability
	"What I Wish I Knew Before Freelancing", // Expert insight
	"The New Rules of Sustainable Travel",   // Modern travel
	"How to Build a Trust-Based Brand",      // Marketing strategy
	"Is the Metaverse Finally Changing?",    // Tech analysis
	"My Secret to 30-Minute Meal Prep",      // Productivity/Food
	"Why Everyone Feels Overstimulated",     // Psychological trend
	"7 Signs You Need a Career Pivot",       // Career advice
	"How to Invest with Inflation in Mind",  // Economic relevance
	"The Case for 'Soft' Productivity",      // Alternative work habits
	"Biohacking Tips for Busy People",       // Health optimization
	"2026 Trends: What's Next for Retail",   // Industry prediction
}

var blogTags = []string{
	"ai-agents",      // High-tech automation
	"wellness",       // Personal health
	"sustainability", // Green initiatives
	"remote-life",    // Modern work culture
	"crypto2026",     // Finance/Web3
	"biohacking",     // Health optimization
	"minimalism",     // Lifestyle aesthetic
	"automation",     // Workflow tools
	"solar-power",    // Energy trends
	"cyber-sec",      // Infrastructure
	"creative-flow",  // Artistic process
	"data-privacy",   // Regulatory focus
	"smart-home",     // IoT/Living
	"mental-health",  // Social awareness
	"open-source",    // Software development
	"travel-tips",    // Leisure
	"future-proof",   // Career/Strategy
	"zero-waste",     // Eco-living
	"gen-alpha",      // Demographic trends
	"productivity",   // Optimization
}

var blogComments = []string{
	"This is exactly what I needed to read today, thanks for sharing!",
	"I tried the AI workflow you mentioned and it saved me hours.",
	"Could you go more into detail about the third point in your next post?",
	"I disagree with the take on remote work, but I appreciate the perspective.",
	"Finally, someone is talking about the ethical side of this tech!",
	"Bookmarking this for later. Great insights for 2026.",
	"Does this strategy work for small teams or just solo creators?",
	"I’ve been following your blog for a while and this is your best one yet.",
	"The tip about digital minimalism really hit home for me.",
	"Can you recommend any specific tools to help automate this process?",
	"Great read! I shared this with my LinkedIn network.",
	"I'm still a bit skeptical, but you've given me a lot to think about.",
	"This is a game changer for my morning routine.",
	"Wait, did you use a specific prompt for the results shown in the header?",
	"I've had a similar experience with biohacking lately—it's transformative.",
	"Short, sweet, and to the point. Love the new formatting.",
	"Are there any updates on how this looks for the Q3 2026 forecast?",
	"I’d love to see a video version of this tutorial!",
	"Thanks for the honest take; most people are just hype-posting right now.",
	"Following! Looking forward to the next update on this series.",
}
var blogContents = []string{
	"Why 2026 is the year we stop 'hustling' and start 'soft' productivity.",
	"A step-by-step breakdown of my daily 15-minute AI-assisted morning routine.",
	"How to build a personal brand in 2026 that doesn't feel like a performance.",
	"I tried digital minimalism for 30 days: Here is what I deleted for good.",
	"The best affordable biohacking tips that actually have scientific backing.",
	"Why your remote work boundaries are failing and how to reset them today.",
	"A guide to 'micro-rest' practices that prevent burnout in a busy world.",
	"How I used AI to automate my home's energy consumption and saved 30%.",
	"The truth about making passive income with digital products in 2026.",
	"Why everyone is talking about 'nervous system regulation' right now.",
	"5 ways to use AI as a creative partner rather than a replacement.",
	"How to design an intentional living space that boosts your mental clarity.",
	"What I wish I knew before I started my career pivot in my 30s.",
	"The top 5 sustainable travel destinations for budget-conscious explorers.",
	"How to navigate the 'death of proof' in an era of deepfakes and AI.",
	"A beginner's guide to investing in green technology and circular economies.",
	"Why low-impact fitness is trending more than high-intensity in 2026.",
	"How to curate your digital feed to prioritize mental health over doomscrolling.",
	"The secrets of 'quiet success': Building a business without social fame.",
	"How to stay human-centric in a world increasingly run by algorithms.",
}

// Seed populates the database with sample users, posts, and comments.
func Seed(st store.Storage, db *sql.DB) {
	ctx := context.Background()

	allUsers := generateUsers(100)
	created_users := make([]*store.User, 0, len(allUsers))
	tx, _ := db.BeginTx(ctx, nil)

	for _, u := range allUsers {
		if err := st.Users.Create(ctx, tx, u); err != nil {
			log.Println("error creating user:", err)
			continue // skip bad users
		}
		created_users = append(created_users, u)
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	posts := generatePosts(100, created_users, rng)
	for _, p := range posts {
		if err := st.Posts.Create(ctx, p); err != nil {
			log.Println("error creating post:", err)
			continue
		}
	}

	comments := generateComments(500, created_users, posts, rng)
	for _, c := range comments {
		if err := st.Comments.Create(ctx, c); err != nil {
			log.Println("error creating comment:", err)
			continue
		}
	}
}

// generateUsers builds a slice of synthetic users.
func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
		if err := users[i].Password.Set("password123"); err != nil {
			continue
		}
	}
	return users
}

// generatePosts builds sample posts tied to the provided users.
func generatePosts(num int, users []*store.User, rng *rand.Rand) []*store.Post {
	posts := make([]*store.Post, 0, num)
	if len(users) == 0 {
		return posts
	}
	for i := 0; i < num; i++ {
		user := users[rng.Intn(len(users))]
		posts = append(posts, &store.Post{
			UserID:  user.ID,
			Title:   blogTitles[rng.Intn(len(blogTitles))],
			Content: blogContents[rng.Intn(len(blogContents))],
			Tags: []string{
				blogTags[rng.Intn(len(blogTags))],
				blogTags[rng.Intn(len(blogTags))],
			},
		})
	}
	return posts
}

// generateComments builds sample comments tied to the provided users and posts.
func generateComments(num int, users []*store.User, posts []*store.Post, rng *rand.Rand) []*store.Comment {
	comments := make([]*store.Comment, 0, num)
	if len(users) == 0 || len(posts) == 0 {
		return comments
	}
	for i := 0; i < num; i++ {
		user := users[rng.Intn(len(users))]
		post := posts[rng.Intn(len(posts))]
		comments = append(comments, &store.Comment{
			UserID:  user.ID,
			PostID:  post.ID,
			Content: blogComments[rng.Intn(len(blogComments))],
		})
	}
	return comments
}
