package mutations

const (
	RegisterUserMutation = `
		mutation($name: String!, $email: String!, $password: String!) {
			insert_users(objects: {name: $name, email: $email, password: $password}) {
				returning {
					id
					name
					email
				}
			}
		}
	`
)
