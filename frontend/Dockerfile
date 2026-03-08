# Build stage
FROM node:22-alpine AS builder

# Enable corepack for pnpm (since Nuxt uses pnpm-lock.yaml)
RUN corepack enable && corepack prepare pnpm@latest --activate

WORKDIR /app

# Copy package.json and pnpm-lock.yaml
COPY package.json pnpm-lock.yaml ./

# Install dependencies
RUN pnpm install --frozen-lockfile

# Copy the rest of the application
COPY . .

# Build the Nuxt application
RUN pnpm build

# Final stage
FROM node:22-alpine AS runner

WORKDIR /app

# Copy the build output from the builder stage
# Nuxt builds into .output by default
COPY --from=builder /app/.output ./.output

# Expose the default Nuxt port
EXPOSE 3000

# Set environment variables
ENV HOST=0.0.0.0
ENV PORT=3000
ENV NODE_ENV=production

# Run the Nitro server
CMD ["node", ".output/server/index.mjs"]
