package queries

const (
	LoginUserQuery = `
		query($email: String!) {
			users(where: {email: {_eq: $email}}) {
				id
				name
				email
				password
				role
			}
		}
	`
)
